var $ = require("jquery"),
    conn = new WebSocket("ws://localhost:7428/ws"),
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
    conn.send("Connected");
};

conn.onmessage = function(e) {
    var message = JSON.parse(e.data);
    // console.log(message);

    switch (message.Category) {
        case "CSS":
            changeStylesheet(message.File);
            break;
        default:
    }
};
