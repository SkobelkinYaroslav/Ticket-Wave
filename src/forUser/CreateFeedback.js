import { useState } from "react";
import { useLocation } from 'react-router-dom';
import axios from 'axios';
import '../App.css';

const CreateFeedback = () => {
    const [name, setName] = useState('');
    const [nameDirty, setNameDirty] = useState(false);
    const [nameError, setNameError] = useState('Поле не может быть пустым');

    const NameHandler = (e) => {
        setName(e.target.value)

        if (!e.target.value) {
            setNameError('Поле не может быть пустым')
        } else {
            setNameError("")
        }
    }

    const blurHandler = (e) => {
        switch (e.target.name) {
            case 'name':
                setNameDirty(true)
                break
            default:
                break
        }
    }

    const location = useLocation();

    const getEventIdFromUrl = () => {
        const query = new URLSearchParams(location.search);
        return query.get('id');
    };

    const handleSubmit = async (e) => {
        e.preventDefault();

        const eventId = getEventIdFromUrl();



        const feedback = {
            eventId: parseInt(eventId),
            text: name
        };

        axios.post('http://localhost:8080/feedback', feedback, {
            withCredentials: true
        })
            .then(function (response) {
                console.log(response.data);
                window.location.href = "/user-profile";
            })
            .catch(function (error) {
                console.error('Error sending feedback:', error);
            });
    };

    return (
        <div className='reg'>
            <h1>Отзыв</h1>
            {(nameDirty && nameError) && <div style={{ color: 'red' }}>{nameError}</div>}
            <div className="input-box-feedback">
                <input onChange={e => NameHandler(e)} value={name} onBlur={e => blurHandler(e)} name='name' type="text" placeholder="Текст Вашего отзыва..." />
            </div>
            <div>
                <button className="register-link" onClick={handleSubmit}>
                    <p><a href="/user-profile">Отправить отзыв</a></p>
                </button>
            </div>
        </div>
    );
};

export default CreateFeedback;