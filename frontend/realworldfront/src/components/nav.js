import React, { useState } from 'react';
import NavbarAuth from './navbarAuth';
import NavbarNoAuth from './navbarNoAuth';




export default function Nav({ isAuthenticated }) {
    return (
        isAuthenticated ? (<NavbarAuth />) : (<NavbarNoAuth />)
    )
};
