var $ = require("jquery"),
    conn = new WebSocket("ws://localhost:8080/ws"),
    changeStylesheet = function(stylesheet) {
        var element,
            override = document.head.querySelector("#sherbet-style");

        if (override) {
            override.remove();
        }

        element = document.createElement("style");
        element.innerHTML = stylesheet;
        element.id = "sherbet-style";
        document.head.appendChild(element);
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
        default:
    }
};
