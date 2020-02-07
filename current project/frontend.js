import React, { Component } from "react";
import axios from "axios";
import { Card, Header, Form, Input, Icon } from "semantic-ui-react";

let endpoint = "http://localhost:8080"

const CardItems = () => (	
	//<Segment inverted> dark background not working yet
		<Card.Group itmesPerRow={1}>
			<Card color='purple'> 
				<Card.Content>
					/* <Image floated='right'/> */
					<Card.Header>{events.Title}</Card.Header> // item name 
					//<Card.Meta></Card.Meta> // brand name
					<Card.Description>{events.Description}</Card.Description> //this will be the item Description
				</Card.Content>
				<Card.Content extra>
					<div className='ui three buttons'>
						<Button inverted color='blue'>
							<Icon name='star' />
	        				have it
	      				</Button>
	      				<Button inverted color='pink' icon>
	      					<Icon name='heart' />
	        				want it
	      				</Button>
	      				<Button inverted color="violet" icon>
	      					<Icon name='fork' />
	      					more info
	      				</Button>
	      			</div>
				</Card.Content>
			</Card>

    	</Card.Group>
    //</Segment>
);

function Frontend(){
	/* should the componentDidmount go here?
	and the CardItem goes here*/
} 

componentDidMount() {
	fetch(endpoint + "/events/1")
	.then(res => res.json())
	.then((data) => {
		this.setState({ events: data})
	})
	.catch(console.log)
}

ReactDOM.render(
	<Frontend>

	</Frontend>
);