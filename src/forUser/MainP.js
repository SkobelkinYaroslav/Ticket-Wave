import React from 'react';
import axios from 'axios';
import '../App.css';
import MenuUser from './MenuUser';
import EventsMainP from './EventsMainP';
import Categories from './Categories';
import ShowFullEvent from './ShowFullEvent';

class MainP extends React.Component {
    state = {
        selectedOption: null,
        currentEvents: [],
        events: [],
        showFullEvent: false,
        fullEvent: {},
    };

    constructor(props) {
        super(props);
        this.chooseCategory = this.chooseCategory.bind(this);
        this.onShowEvent = this.onShowEvent.bind(this);
    }

    componentDidMount() {
        axios.get('http://localhost:8080/event', { withCredentials: true })
            .then((response) => {
                const events = response.data.events.map(event => ({
                    id: event.id,
                    OrganizerID: event.organizerId,
                    Name: event.name,
                    category: event.category,
                    Description: event.description,
                    DateTime: event.dateTime,
                    Address: event.address,
                    TicketPrice: event.ticketPrice,
                    TicketCount: event.ticketCount,
                    img: event.img,
                }));
                this.setState({ events, currentEvents: events });
            })
            .catch((error) => {
                console.error('Error fetching events:', error);
            });
    }

    handleChange = (selectedOption) => {
        this.setState({ selectedOption }, () =>
            console.log(`Option selected:`, this.state.selectedOption)
        );
    };

    chooseCategory(category) {
        if (category === 'all') {
            this.setState({ currentEvents: this.state.events });
            return;
        }
        this.setState({
            currentEvents: this.state.events.filter((el) => el.category === category),
        });
    }

    onShowEvent(event) {
        this.setState({ fullEvent: event, showFullEvent: !this.state.showFullEvent });
    }

    render() {
        const { selectedOption } = this.state;

        return (
            <form>
                <div className="search">
                    <a className="title">TicketWave - Билеты на мероприятия</a>
                    <MenuUser />
                </div>
                <Categories chooseCategory={this.chooseCategory} />
                <EventsMainP onShowEvent={this.onShowEvent} events={this.state.currentEvents} />
                {this.state.showFullEvent && <ShowFullEvent onShowEvent={this.onShowEvent} event={this.state.fullEvent} />}
            </form>
        );
    }
}

export default MainP;
