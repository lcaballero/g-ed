import React from "react";
import ReactDOM from "react-dom";
import Start from "./start-websocket.js";

require("file-loader?name=index.html!../index.html");
require("../css/reset.css");
require("../css/render.css");

Start();

class InputLine extends React.Component {
    render() {
        return <textarea id="input-line" cols="85" rows="3" defaultValue=" ">
        </textarea>
    }
}

class Send extends React.Component {
    render() {
        return <button class="send" type="button">Send</button>
    }
}

class Posts extends React.Component {
    render() {
        return <div class="posts"> </div>
    }
}

class Inputs extends React.Component {
    render() {
        return <div class="inputs">
            <InputLine/>
            <Send/>
        </div>
    }
}

class Layout extends React.Component {
	render() {
		const { title } = this.props;
		return (
			<div class="main">
                <h1>{title}</h1>
                <Posts/>
                <Inputs/>
            </div>
		)
	}
}

const render = () => {
    const app = document.getElementById('app');
    ReactDOM.render(<Layout title="Chatting" />, app);
};

render();
