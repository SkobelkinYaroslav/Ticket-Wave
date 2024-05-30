import React from 'react';
import axios from 'axios';
import EventsProfile from "./EventsProfile";

class Profile extends React.Component {
    state = {
        selectedOption: null,
        events: [],
    };

    constructor(props) {
        super(props);
        this.handleChange = this.handleChange.bind(this);
    }

    componentDidMount() {
        axios.get('http://localhost:8080/user_events', { withCredentials: true })
            .then(response => {
                const events = response.data.events.map(event => ({
                    id: event.id,
                    OrganizerID: event.organizerId,
                    Name: event.name,
                    category: event.category,
                    Description: event.description,
                    DateTime: event.dateTime,
                    Venue: event.venue,
                    Address: event.address,
                    TicketPrice: event.ticketPrice,
                    TicketCount: event.ticketCount,
                    img: event.img,
                }));
                this.setState({ events });
            })
            .catch(error => {
                console.error('Error fetching user events:', error);
            });
    }

    handleChange(selectedOption) {
        this.setState({ selectedOption }, () =>
            console.log(`Option selected:`, this.state.selectedOption)
        );
    }

    render() {
        const { selectedOption, events } = this.state;

        return (
            <div>
                <div className="search">
                    <h1>Профиль пользователя</h1>
                </div>
                <div className="text">
                    <h2>Мои билеты</h2>
                </div>
                <EventsProfile events={events} />
            </div>
        );
    }
}

export default Profile;
