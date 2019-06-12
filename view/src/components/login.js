import React  from 'react';

const Login  = () =>{
    return(
        <div>
            <h1>This is a login screen</h1>
            <form>
                <label>
                    Nome de usuário:
                    <input type="text" name="username"/>
                </label>
                <br/>
                <label>
                    Senha:
                    <input type="text" name="password"/>
                </label>
                <br/>
                <input type="submit" value="submit"/>
            </form>
        </div>
    )
}

export default Login
