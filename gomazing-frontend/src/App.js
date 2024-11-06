import React from 'react';
import { BrowserRouter as Router, Route, Routes, Link } from 'react-router-dom';
import MazeApp from './MazeApp';
import FloodFillApp from './FloodFillApp';
import './App.css';

function App() {
    return (
        <Router>
            <div className="App">
                <nav>
                    <ul>
                        <li>
                            <Link to="/">Maze</Link>
                        </li>
                        <li>
                            <Link to="/floodfill">Flood Fill</Link>
                        </li>
                    </ul>
                </nav>
                <Routes>
                    <Route path="/floodfill" element={<FloodFillApp />} />
                    <Route path="/" element={<MazeApp />} />
                </Routes>
            </div>
        </Router>
    );
}

export default App;