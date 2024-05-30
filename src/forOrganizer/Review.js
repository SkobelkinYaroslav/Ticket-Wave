import React, { Component } from "react";

export class Review extends Component{
    render(){
        return(
            <div className="review">

                <h2>{this.props.review.NameUser}</h2>
                <p>{this.props.review.textReview}</p>
                
            </div>
        )
    }
}
export default Review

