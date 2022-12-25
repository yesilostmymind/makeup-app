import React from 'react';
import './App.css';
import { Container } from "semantic-ui-react";
import MakeupList from "./app-list";

function App() {
  return (
    <div>
      <Container>
        <MakeupList />
      </Container>
    </div>
  );
}
export default App;
