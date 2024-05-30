import React, { Component } from "react";
import axios from 'axios';

export class EventMainP extends Component {
    handleBuyTicket = () => {
        const userEventLink = {
            id: this.props.event.id,
            event: {
                id: this.props.event.id,
            },
            linkType: "buy"
        };

        axios.post('http://localhost:8080/buy', userEventLink, {
            withCredentials: true
        })
            .then(response => {
                window.location.href = "/user-profile";
            })
            .catch(error => {
                console.error('There was an error purchasing the ticket:', error);
            });
    }

    render() {
        return (
            <div className="event">
                <img src={"./img/" + this.props.event.img} onClick={() => this.props.onShowEvent(this.props.event)} alt="Image" />
                <h2>{this.props.event.Name}</h2>
                <p>{this.props.event.DateTime}</p>
                <div className="add-to-cart" onClick={this.handleBuyTicket}><a href="/user-profile">Купить билет</a></div>
                {/* onClick={() => this.props.onAdd(this.props.event)} */}
                {/* <div className="register-link"><p><a href="/create-feedback">Написать отзыв</a></p></div> */}
            </div>
        );
    }
}

export default EventMainP;
