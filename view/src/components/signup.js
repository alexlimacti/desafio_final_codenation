import React  from 'react';

const Signup  = () =>{
    return(
        <div>
            <h1>This is a sign up screen</h1>
            <form>
                <label>
                    Nome de usuário:
                    <input type="text" name="username"/>
                </label>
                <br/>
                <label>
                    Escolha uma senha:
                    <input type="text" name="password"/>
                </label>
                <br/>
                <label>
                    Repita a senha:
                    <input type="text" name="re-password"/>
                </label>
                <br/>
                <input type="submit" value="submit"/>
            </form>
        </div>
    )
}

export default Signup