import React from 'react';
import ReactDOM from 'react-dom';
import "./App.css";

import { Container, Header, List } from "semantic-ui-react";

import Frontend from "./frontend";

const App = ({ children }) => (
	<Container>
		<Segment inverted>
			<Header as="h1" color="violet">MakeupDB</Header>
			<p>this is an app to store your makeup</p>
		</Segment inverted>
		
		{children}
	</Container>
);


const styleLink = document.createElement("link");
styleLink.rel = "stylesheet";
styleLink.href = "https://cdn.jsdelivr.net/npm/semantic-ui/dist/semantic.min.css";
document.head.appendChil(styleLink);

ReactDOM.render(
	<App>
		<Frontend />
	</App>,
	document.getElementById('app') //is this right?
);
 
