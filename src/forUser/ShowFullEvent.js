import { render } from "@testing-library/react";
import React, { Component } from "react";
import axios from "axios";

export class ShowFullEvent extends Component{
    handleBuyTicket = (e) => {
        e.preventDefault(); // предотвратить переход по ссылке

        const userEventLink = {
            event: {
                id: this.props.event.id,
            },
            linkType: "buy"
        };

        axios.post('http://localhost:8080/buy', userEventLink, {withCredentials: true})
            .then(response => {
                window.location.href = "/user-profile";
            })
            .catch(error => {
                console.error('There was an error purchasing the ticket:', error);
            });
    }

    render(){
        return(
            <div className="full-event">
                <div className="event">
                    <div>
                        <img src={"./img/"  + this.props.event.img } onClick={()=> this.props.onShowEvent(this.props.event)} alt="Image"/>
                        <h1>{this.props.event.Name}</h1>
                        <div>{this.props.event.Description}</div>
                        <p>Дата проведения: {this.props.event.DateTime}</p>
                        <p>Место проведения: {this.props.event.Address}</p>
                        <p>Цена билета: {this.props.event.TicketPrice} руб.</p>

                        <button className="add-to-cart" onClick={this.handleBuyTicket}>
                            Купить билет
                        </button>

                    </div>
                </div>
            </div>
        )
    }
}

export default ShowFullEvent;
