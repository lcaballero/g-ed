import React from "react";
import ReactDOM from "react-dom";
import Start from "./start-websocket";
import { store } from "./store";

require("file-loader?name=index.html!../index.html");
require("../css/reset.css");
require("../css/render.css");

Start();

// this is bad but simple
var inputArea = null;

class InputLine extends React.Component {
    constructor(props) {
        super(props);

        this.handleChange = this.handleChange.bind(this);
    }
    handleChange(ev) {
        var message = ev.target.value;
        inputArea = ev.target;
        store.dispatch({type:"current-message", value: message});
        this.setState({value: message})
    }
    render() {
        return <textarea
            refs="messageInput"
            id="input-line" cols="85" rows="3" defaultValue=" "
            onChange={this.handleChange}>
        </textarea>
    }
}

class Send extends React.Component {
    handle(ev) {
        var message = store.getState().currentMessage || "";
        console.log("handle", ev, message);
        store.dispatch({type: "ADD_POST", message: message})
    }
    render() {
        return <input type="submit" class="send" onClick={this.handle} value="Send"/>
    }
}

class Posts extends React.Component {
    render() {
        const { posts } = store.getState();
        var els = [];
        for (var i = 0; i < posts.length; i++) {
            var { message } = posts[i];
            els.push(<div key={i}>{message}</div>)
        }

        return <div class="posts">
            {els}
        </div>
    }
}

class Inputs extends React.Component {
    constructor(props) {
        super(props);

        this.handleSubmit = this.handleSubmit.bind(this);
    }
    handleSubmit(ev) {
        ev.preventDefault();
        var area = inputArea || {};
        area.value = "";
    }
    render() {
        return <form class="inputs" onSubmit={this.handleSubmit}>
            <InputLine/>
            <Send/>
        </form>
    }
}

class Layout extends React.Component {
	render() {
		const { title } = this.props;
		return (
			<div>
                <h1>{title}</h1>
                <Posts/>
                <Inputs/>
            </div>
		)
	}
}

class Main extends React.Component {
    render() {

        return (
            <div class="main">
                <Layout title="Chatting"/>
            </div>
        )
    }
}

const render = () => {
    const app = document.getElementById('app');
    ReactDOM.render(<Main />, app);
};

store.subscribe(render);
render();