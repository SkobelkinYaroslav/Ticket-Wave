import React, { Component } from "react";

export class EventProfile extends Component{
    render(){
        return(
            <div className="event">
                <img src={"./img/"  + this.props.event.img } alt="Image"/>
                <h2>{this.props.event.Name}</h2>
                <p>{this.props.event.DateTime}</p>
                <div className="register-link"><p><a href={`/create-feedback?id=${this.props.event.id}`}>Написать отзыв</a></p></div>
            </div>
            
        )
    }
}
export default EventProfile