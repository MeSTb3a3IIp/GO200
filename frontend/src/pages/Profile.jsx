import React from "react";
import "../css/Profile.css"


export default function Profile () {
    return (
        <div className="bg">
            <div className="design">
                <div className="navbar"></div>
                <div className="container">
                    <div className="main">
                        <div className="self">
                            <div className="avatar"></div>
                            <div className="username"></div>
                            <div className="message"></div>
                        </div>
                        <div className="scores">
                            <div className="stars"></div>
                            <div className="easy"></div>
                            <div className="medium"></div>
                            <div className="hard"></div>
                        </div>
                        <div className="notifications">
                        {/*notify notify notify*/}
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
}