import React from 'react';
import { BrowserRouter as Router, Route, Routes, Link } from 'react-router-dom';
import MazeApp from './MazeApp';
import FloodFillApp from './FloodFillApp';
import MazeGeneration from './MazeGeneration';
import './App.css';

function App() {
    return (
        <Router basename={process.env.PUBLIC_URL}>
            <div className="App">
                <nav>
                    <ul>
                        <li>
                            <Link to="/">Maze</Link>
                        </li>
                        <li>
                            <Link to="/floodfill">Flood Fill</Link>
                        </li>
                        <li>
                            <Link to="/new-maze">New Maze</Link>
                        </li>
                    </ul>
                </nav>
                <Routes>
                    <Route path="/floodfill" element={<FloodFillApp />} />
                    <Route path="/new-maze" element={<MazeGeneration />} />
                    <Route path="/" element={<MazeApp />} />
                </Routes>
            </div>
        </Router>
    );
}

export default App;