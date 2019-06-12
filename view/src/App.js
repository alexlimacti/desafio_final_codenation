import React, { Component } from 'react';
import { Switch, Route } from "react-router-dom";
import Login from "./components/login";
import Signup from "./components/signup";
import Dashboard from "./components/dashboard";
import AlertList from "./components/alert_list"
import './App.css';

class App extends Component {
  render(){
    return (
      <div className="App">
        <Switch>
          <Route path="/" exact component={Login} />
          <Route path="/signup" component={Signup} />
          <Route path="/dashboard" component={Dashboard} />
          <Route path="/alert-list" component={AlertList} />
        </Switch>
      </div>
    );
  }
  
}

export default App;
