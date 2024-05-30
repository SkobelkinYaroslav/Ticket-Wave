import React, { useState, useEffect } from "react";
import axios from "axios";
import Reviews from "./Reviews";
import { useLocation } from 'react-router-dom';

const Feedback = () => {
    const [reviews, setReviews] = useState([]);
    const location = useLocation();

    const getEventIdFromUrl = () => {
        const query = new URLSearchParams(location.search);
        return query.get('id');
    };

    useEffect(() => {
        const eventId = getEventIdFromUrl();
        console.log(eventId);

        axios.get(`http://localhost:8080/feedback?id=${eventId}`, { withCredentials: true })
            .then(response => {
                console.log('Feedback:', response.data.feedback)
                const reviewsData = response.data.feedback.map(feedback => ({
                    id: feedback.id,
                    NameUser:feedback.senderName,
                    textReview: feedback.text,
                }));

                setReviews(reviewsData);
                console.log("reviewData ", reviewsData)
                console.log("review ", reviews);
            })
            .catch(error => {
                console.error('Error fetching feedback:', error);
            });
    }, [location]);

    return (
        <div>
            <div className="search"><h1>Отзывы на ваше мероприятие</h1></div>
            <Reviews review={reviews} />
        </div>
    );
};

export default Feedback;
