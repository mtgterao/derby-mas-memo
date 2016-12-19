require('material-design-lite/dist/material-grid.min.css');
require('material-design-lite/dist/material.blue_grey-light_blue.min.css');
require('material-design-lite/dist/material.min.js');

import React from 'react';
import ReactDOM from 'react-dom';
import 'react-mdl/extra/material.css';
import 'react-mdl/extra/material.js';

import AppLayout from './components/app_layout.jsx';

(function() {
    ReactDOM.render(
        <AppLayout />,
        document.getElementById('main')
    );
})();
