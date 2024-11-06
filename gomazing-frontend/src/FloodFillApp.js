import React, { useState, useEffect } from 'react';
import axios from 'axios';
import NodeComponent from './NodeComponent';
import './FloodFillApp.css';

const FloodFillApp = () => {
    const [width, setWidth] = useState(Math.trunc(window.innerWidth / 20));
    const [height, setHeight] = useState(Math.trunc(window.innerHeight / 20));
    const [grid, setGrid] = useState([]);
    const [maze, setMaze] = useState([]);
    const nodeSize = 20;

    useEffect(() => {
        initializeGrid();
    }, [width, height]);

    useEffect(() => {
        const handleResize = () => {
            setWidth(Math.trunc(window.innerWidth / nodeSize));
            setHeight(Math.trunc(window.innerHeight / nodeSize));
        };

        window.addEventListener('resize', handleResize);
        return () => window.removeEventListener('resize', handleResize);
    }, []);

    const initializeGrid = () => {
        const newGrid = Array.from({ length: height }, () =>
            Array.from({ length: width }, () => ({
                filled: false,
                walls: [true, true, true, true],
            }))
        );
        setGrid(newGrid);
    };

    const generateMaze = async (algo) => {
        try {
            const response = await axios.post('http://localhost:8080/generate', {
                width,
                height,
                algo
            });
            setMaze(response.data);
        } catch (error) {
            console.error('Error generating maze:', error);
        }
    };

    const handleClick = (rowIndex, cellIndex) => {
        floodFill(rowIndex, cellIndex);
    };

    const floodFill = (x, y) => {
        if (x < 0 || y < 0 || x >= height || y >= width || grid[x][y].filled) {
            return;
        }

        const newGrid = [...grid];
        newGrid[x][y].filled = true;
        setGrid(newGrid);

        setTimeout(() => {
            floodFill(x - 1, y);
            floodFill(x + 1, y);
            floodFill(x, y - 1);
            floodFill(x, y + 1);
        }, 100);
    };

    return (
        <div className="floodfill-app">
            <h1>Flood Fill Algorithm with Maze Generation</h1>
            <div className="controls">
                <button onClick={() => generateMaze(1)}>Generate Maze with DFS</button>
                <button onClick={() => generateMaze(2)}>Generate Maze with Prim's</button>
                <button onClick={() => generateMaze(3)}>Generate Maze with Kruskal's</button>
                <button onClick={() => generateMaze(4)}>Generate Maze with BFS</button>
            </div>
            <div className="floodfill-grid" style={{ gridTemplateColumns: `repeat(${width}, ${nodeSize}px)` }}>
                {grid.map((row, rowIndex) =>
                    row.map((cell, cellIndex) => {
                        const mazeCell = maze[rowIndex] ? maze[rowIndex][cellIndex] : null;
                        return (
                            <NodeComponent
                                key={`${rowIndex}-${cellIndex}`}
                                nodeNumber={rowIndex * width + cellIndex}
                                size={nodeSize}
                                walls={mazeCell ? mazeCell.walls : [true, true, true, true]}
                                isActive={cell.filled}
                                onClick={() => handleClick(rowIndex, cellIndex)}
                            />
                        );
                    })
                )}
            </div>
        </div>
    );
};

export default FloodFillApp;