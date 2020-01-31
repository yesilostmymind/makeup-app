import React, { Component } from "react";
import axios from "axios";
import { Card, Header, Form, Input, Icon } from "semantic-ui-react";

const CardItems = () (
	//<Segment inverted> dark background not working yet
		<Card.Group itmesPerRow={1}>
			<Card color='purple'> 
				<Card.Content>
					<Image floated='right'/>
					<Card.Header> </Card.Header> //this will be the item name
					<Card.Description> </Card.Description> //this will be the item Description
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
	    	<Card color='purple'>
		    	<Card.Content>
						<Card.Header> </Card.Header>
						<Card.Description> </Card.Description>
				</Card.Content>
				<Card.Content extra>
					<div className='ui three buttons'>
						<Button inverted color='blue' icon>
							<Icon name='heart' />
	        				have it
	      				</Button>
	      				<Button inverted color='pink' icon>
	      					<Icon name='star' />
	        				want it
	      				</Button>
	      				<Button inverted color="violet" icon>
	      					<Icon name='fork' />
	      					more info
	      				</Button>
      				</div>
				</Card.Content>
	    	</Card>
	    	<Card color='purple'>
		    	<Card.Content>
					<Card.Header> </Card.Header>
					<Card.Description> </Card.Description>
				</Card.Content>
				<Card.Content extra>
					<div className='ui three buttons'>
					<Button inverted color='blue'>
		        		Have it
		      		</Button>
		      		<Button inverted color='pink'>
		        		want it
		      		</Button>
		      		<Button inverted color="violet">
		      			more info
		      		</Button>
				</Card.Content>
	    	</Card>
	    	<Card color='purple'>
	    		<Card.Content>
					<Card.Header> </Card.Header>
					<Card.Description> </Card.Description>
				</Card.Content>
				<Card.Content extra>
					<Button inverted color='blue'>
        				Have it
      				</Button>
      				<Button inverted color='pink'>
        				want it
      				</Button>
      				<Button inverted color="violet">
      					more info
      				</Button>
				</Card.Content>
	    	</Card>
	    	<Card color='purple'>
	    		<Card.Content>
					<Card.Header> </Card.Header>
					<Card.Description> </Card.Description>
				</Card.Content>
				<Card.Content extra>
					<Button inverted color='blue'>
        				Have it
      				</Button>
      				<Button inverted color='pink'>
        				want it
      				</Button>
      				<Button inverted color="violet">
      					more info
      				</Button>
				</Card.Content>
	    	</Card>
    	</Card.Group>
    //</Segment>
)

/*
let endpoint = "http://localhost:8080";

class MakeupDataBase extends Component {
	constructor(props) {
		super(props);

		this.state = {
			task: "",
			items: []
		};
	}

	compoenntDidMount() {
		this.getTask();
	}
}
*/