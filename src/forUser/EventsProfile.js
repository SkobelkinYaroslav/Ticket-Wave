import React from "react";
import EventProfile from './EventProfile'
import '../App.css';
export class EventsProfile extends React.Component{
    render(){
        return(
            <main>
                {this.props.events.map(el=>(
                    <EventProfile key={el.id} event={el}  />
                ))}
            </main>
        )
    }
}
 export default EventsProfile 