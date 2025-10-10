import React from "react";
import "./css/Stats.css"

export default function Stats () {
    return (
        <div className="bg">
            <div className="design">
                <div className="navbar"></div>
                <div className="container">
                    <div className="main">
                        <div className="title"></div>
                        <div className="speed">
                            <div className="graph"></div>
                            <div className="description"></div>
                        </div>
                        <div className="memory">
                            <div className="graph"></div>
                            <div className="description"></div>
                        </div>
                        <div className="result">
                            <div className="description"></div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
}