import React, { Component } from "react";

export class EventOrganizer extends Component{
    render(){
        return(
            <div className="event">
                <img src={"./img/"  + this.props.event.img } alt="Image"/>
                <h2>{this.props.event.Name}</h2>
                <p>{this.props.event.DateTime}</p>
                <div className="register-link"><p><a href={"/feedback?id=" + this.props.event.id}>Просмотреть отзывы</a>
                </p></div>
            </div>
        )
    }
}

export default EventOrganizer