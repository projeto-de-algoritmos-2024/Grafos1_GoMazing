import React, { useState } from 'react';
import axios from 'axios';
import './MazeApp.css';

const MazeApp = () => {
    const [width, setWidth] = useState(10);
    const [height, setHeight] = useState(10);
    const [maze, setMaze] = useState(null);
    const [steps, setSteps] = useState([]);
    const [currentStep, setCurrentStep] = useState(0);

    const generateMaze = async (algo) => {
        try {
            const response = await axios.post('http://localhost:8080/generate', {
                width,
                height,
                algo
            });
            setSteps(response.data);
            setCurrentStep(0);
            setMaze(response.data[0]);
            animateMaze(response.data);
        } catch (error) {
            console.error('Error generating maze:', error);
        }
    };

    const animateMaze = (steps) => {
        let stepIndex = 0;
        const interval = setInterval(() => {
            if (stepIndex < steps.length) {
                setMaze(steps[stepIndex]);
                setCurrentStep(stepIndex); // Update current step
                stepIndex++;
            } else {
                clearInterval(interval);
            }
        }, 1500 / steps.length); // Adjust the speed as needed
    };

    const solveMaze = async (solveAlgo) => {
        try {
            const response = await axios.post('http://localhost:8080/solve', {
                algo: solveAlgo
            });
            setMaze(response.data);
        } catch (error) {
            console.error('Error solving maze:', error);
        }
    };

    return (
        <div className="maze-app">
            <h1>GoMazing</h1>
            <div className="controls">
                <label>
                    Width:
                    <input type="number" value={width} onChange={(e) => setWidth(e.target.value)} />
                </label>
                <label>
                    Height:
                    <input type="number" value={height} onChange={(e) => setHeight(e.target.value)} />
                </label>
            </div>
            <div className="buttons">
                <button onClick={() => generateMaze(1)}>Generate with DFS</button>
                <button onClick={() => generateMaze(2)}>Generate with Prim's</button>
                <button onClick={() => generateMaze(3)}>Generate with Kruskal's</button>
                <button onClick={() => generateMaze(4)}>Generate with BFS</button>
            </div>
            {maze && (
                <div className="maze-display">
                    <h2>Maze</h2>
                    <pre>{JSON.stringify(maze, null, 2)}</pre>
                    <button onClick={() => solveMaze(1)}>Solve with DFS</button>
                    <button onClick={() => solveMaze(2)}>Solve with BFS</button>
                </div>
            )}
        </div>
    );
};

export default MazeApp;