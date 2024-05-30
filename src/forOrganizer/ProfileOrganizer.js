import React from "react";
import EventsOrganizer from './EventsOrganizer'
import axios from "axios";
class ProfileOrganizer extends React.Component{
    state = {
        selectedOption: null,
      };
      handleChange = (selectedOption) => {
        this.setState({ selectedOption }, () =>
          console.log(`Option selected:`, this.state.selectedOption)
        );
      };
      constructor(props){
        super(props)
        this.state ={
            events: [
            ]
        }
    }
    componentDidMount() {
        axios.get('http://localhost:8080/organizer_events', { withCredentials: true })
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

    render() {
        const { selectedOption } = this.state;
    return (
        <div>
   <div className="search">
<h1>Профиль организатора</h1>
<div> <button className="create-event"><a href="/create-event">Создать мероприятие</a></button> </div>
</div>
<div className="text"><h2>Мои мероприятия</h2></div>
<EventsOrganizer events={this.state.events}/></div>
    );
}
}
export default ProfileOrganizer;


