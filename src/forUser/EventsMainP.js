import React from "react";
import EventMainP from './EventMainP'
import '../App.css';
export class EventsMainP extends React.Component{
    render(){
        return(
            <main>
                {this.props.events.map(el=>(
                    <EventMainP onShowEvent={this.props.onShowEvent} key={el.id} event={el}  />
                ))}
            </main>
        )
    }
}
 export default EventsMainP 

