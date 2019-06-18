import React, { Component } from 'react';
import { withRouter } from 'react-router';
import PropTypes from 'prop-types';

import loginService from '../services/loginService';

class Login extends Component {
  constructor(props) {
    super(props);
    this.state = {
      username: '',
      password: '',
    };
  }
  handleChange(event) {
    this.setState({ [event.target.name]: event.target.value });
  }
  async login(event) {
    event.preventDefault();

    try {
      await loginService.login(this.state.email, this.state.password);
      this.props.history.push('/');
    } catch (error) {
      event.preventDefault();
      alert(error);
    }
  }

  render = () => (
    <div className="container">
      <br />
      <div className="text-center mb-4">
        <h1 className="h3 mb-3 font-weight-normal">Login</h1>
      </div>
      <div id="login-row" class="row justify-content-center align-items-center">
        <div id="login-column" class="col-md-6">
          <div id="login-box" class="col-md-12">
            <form className="form-signin">
              <div className="form-label-group">
                <label htmlFor="inputEmail">Email</label>
                <input
                  name="email"
                  onChange={e => {
                    this.handleChange(e);
                  }}
                  value={this.state.email}
                  className="form-control"
                  placeholder="Email"
                  required
                />
                <div className="form-label-group mt-2">
                  <label htmlFor="inputPassword">Password</label>
                  <input
                    name="password"
                    onChange={e => {
                      this.handleChange(e);
                    }}
                    value={this.state.password}
                    type="password"
                    className="form-control"
                    placeholder="Password"
                    required
                  />
                </div>
                <div className="mt-5">
                  <button
                    type="submit"
                    onClick={e => {
                      this.login(e);
                    }}
                    className="login btn btn-lg btn-primary btn-block"
                  >
                    Login
                  </button>
                </div>
              </div>
            </form>

          </div>
        </div>
      </div>


    </div>
  );
  static propTypes = {
    location: PropTypes.object.isRequired,
    history: PropTypes.object.isRequired,
  };
}

export default withRouter(Login);
