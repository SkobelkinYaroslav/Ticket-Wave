import React from "react";
import '../App.css';
import EventOrganizer from "./EventOrganizer";
export class EventsOrganizer extends React.Component{
    render(){
        return(
            <main>
                {this.props.events.map(el=>(
                    <EventOrganizer key={el.id} event={el}  />
                ))}
            </main>
        )
    }
}
 export default EventsOrganizer