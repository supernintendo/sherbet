var $ = require("jquery"),
    conn = new WebSocket("ws://localhost:8080/ws"),
    adjustFrameHeight = function() {
        $("iframe[name='sherbet-frame']").css({
            height: $(window).height() + "px"
        });
    },
    changeFrame = function(frame) {
        $("iframe[name='sherbet-frame']").attr("src", frame);
    },
    changeStylesheet = function(stylesheet) {
        var element,
            head = window.frames["sherbet-frame"].document.head,
            override = head.querySelector("#sherbet-style");

        if (override) {
            override.remove();
        }

        element = document.createElement("style");
        element.innerHTML = stylesheet;
        element.id = "sherbet-style";
        head.appendChild(element);
    };

conn.onclose = function(e) {
    // TODO: Let server know about closed connections.
};

conn.onopen = function(e) {
    conn.send("connected");
};

conn.onmessage = function(e) {
    var data = e.data.split(":sherbet###");

    switch (data[0]) {
        case "css":
            changeStylesheet(data[1]);
            break;
        case "frame":
            changeFrame(data[1]);
            break;
        default:
    }
};

$(document).ready(function() {
    $(window).on("resize", $.proxy(adjustFrameHeight, this));
    adjustFrameHeight();
});