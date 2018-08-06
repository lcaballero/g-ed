var ws = new WebSocket("ws://localhost:3000/websocket");

export default function Start() {
    ws.onopen = function() {
        // Web Socket is connected, send data using send()
        ws.send(JSON.stringify({
            method:"PUT",
            path:"/join/room-name",
            data:{}
        }));
    };

    ws.onmessage = function (evt) {
        var obj = JSON.parse(evt.data);
        if (obj.type !== "ping") {
            console.log(evt);
        }
    };

    ws.onclose = function () {
        console.log("closing connection")
    };

    window.broadcast = function (obj) {
        ws.send(JSON.stringify(obj))
    };

    // document.getElementById("send").addEventListener("click", function (ev) {
    //     broadcast({
    //         method: "post",
    //         path: "/speak"
    //     })
    // });
}

