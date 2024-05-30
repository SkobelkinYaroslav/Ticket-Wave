import React from "react";
import '../App.css';
import Review from "./Review";
export class Reviews extends React.Component{
    render(){
        return(
            <main>
                {this.props.review.map(el=>(
                    <Review key={el.id} review={el}  />
                ))}
            </main>
        )
    }
}
 export default Reviews;
