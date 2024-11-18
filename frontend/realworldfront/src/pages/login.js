import React, { useState } from 'react';
import PropTypes from 'prop-types';


async function loginUser(credentials) {

    try {
        const response = await fetch("https://localhost:8080/login", {
            method: 'POST',
            headers: { 'Content Type': 'application/json' },
            body: JSON.stringify(credentials)
        });
        console.log("hello")
        if (!response.ok) {
            const error = await response.json()
            console.log("hello2")

            throw new Error(error.message || 'Failed to log in.');
        }
        return await response.json();
    }
    catch (err) {
        console.log("hello3")

        throw new Error("Network Failed")

    }
}

function Login({ setToken }) {
    const [username, setUserName] = useState();
    const [password, setPassword] = useState();
    const [error, setError] = useState(null); // State to store error messages

    const handleSubmit = async e => {
        e.preventDefault();
        setError(null);
        try {
            const token = await loginUser({ username, password });
            setToken(token);
        } catch (err) {
            setError(err.message);
        }
    }



    return (
        <div class="auth-page">
            <div class="container page">
                <div class="row">
                    <div class="col-md-6 offset-md-3 col-xs-12">
                        <h1 class="text-xs-center">Sign in</h1>
                        <p class="text-xs-center">
                            <a href="/register">Need an account?</a>
                        </p>

                        {/* Render error messages */}
                        {error && (
                            <ul className="error-messages">
                                <li>{error}</li>
                            </ul>
                        )}
                        <form onSubmit={handleSubmit}>
                            <fieldset class="form-group">
                                <input type="text" class="form-control form-control-lg" placeholder="Username" onChange={e => setUserName(e.target.value)} />
                            </fieldset>
                            <fieldset class="form-group">
                                <input type="password" class="form-control form-control-lg" placeholder="Password" onChange={e => setPassword(e.target.value)} />
                            </fieldset>
                            <button type="submit" class="btn btn-lg btn-primary pull-xs-right">Sign in</button>
                        </form>
                    </div>
                </div>
            </div>
        </div>
    )

}


Login.propTypes = {
    setToken: PropTypes.func.isRequired
};

export default Login;
