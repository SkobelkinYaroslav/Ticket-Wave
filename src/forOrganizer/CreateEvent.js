import { useState, useRef } from "react";
import '../App.css';
import axios from 'axios';

const CreateEvent = () => {
    const [name, setName] = useState('');
    const [nameCat, setNameCat] = useState('');
    const [description, setDescription] = useState('');
    const [dateTime, setDateTime] = useState('');
    const [address, setAddress] = useState('');
    const [ticketPrice, setTicketPrice] = useState('');
    const [ticketCount, setTicketCount] = useState('');
    const [nameDirty, setNameDirty] = useState(false);
    const [nameCatDirty, setNameCatDirty] = useState(false);
    const [descriptionDirty, setDescriptionDirty] = useState(false);
    const [dateTimeDirty, setDateTimeDirty] = useState(false);
    const [addressDirty, setAddressDirty] = useState(false);
    const [ticketPriceDirty, setTicketPriceDirty] = useState(false);
    const [ticketCountDirty, setTicketCountDirty] = useState(false);
    const [nameError, setNameError] = useState('Поле не может быть пустым');
    const [nameCatError, setNameCatError] = useState('Недопустимая категория');
    const [descriptionError, setDescriptionError] = useState('Поле не может быть пустым');
    const [dateTimeError, setDateTimeError] = useState('Поле не может быть пустым');
    const [addressError, setAddressError] = useState('Поле не может быть пустым');
    const [ticketPriceError, setTicketPriceError] = useState('Поле не может быть пустым');
    const [ticketCountError, setTicketCountError] = useState('Поле не может быть пустым');

    const imageRef = useRef(null);

    const submitHandler = (event) => {
        event.preventDefault();

        const imageFile = imageRef.current.files[0];
        const reader = new FileReader();

        reader.onloadend = () => {
            const base64Image = reader.result.split(',')[1];

            const categoryDictionary = {
                "Концерт": "concert",
                "Спектакль": "performance",
                "Кино": "cinema",
                "Стэндап": "standup"
            };

            const formData = {
                name,
                description,
                category: categoryDictionary[nameCat],
                dateTime,
                address,
                ticketPrice: parseFloat(ticketPrice),
                ticketCount: parseInt(ticketCount),
                img: imageFile.name,
                imgData: base64Image
            };

            axios.post('http://localhost:8080/event', formData, { withCredentials: true })
                .then(response => {
                    window.location.href = '/organizer-profile';
                })
                .catch(error => {
                    console.error('Error creating event:', error);
                });
        };

        reader.readAsDataURL(imageFile);
    };

    const blurHandler = (e) => {
        switch (e.target.name) {
            case 'name':
                if (!e.target.value){
                    setNameDirty(true);
                }else {
                    setNameDirty(false);
                }
                break;
            case 'nameCat':
                switch (e.target.value){
                    case 'Кино':
                    case 'Стэндап':
                    case 'Спектакль':
                    case 'Концерт':
                        setNameCatDirty(false);
                        break;
                    default:
                        setNameCatDirty(true);
                        break;
                }
                break;
            case 'description':
                if (!e.target.value) {
                    setDescriptionDirty(true);
                }else{
                    setDescriptionDirty(false);
                }
                break;
            case 'dateTime':
                if (!e.target.value) {
                    setDateTimeDirty(true);
                }else {
                    setDateTimeDirty(false);
                }
                break;
            case 'address':
                if (!e.target.value){
                    setDateTimeDirty(true)
                }else {
                    setDateTimeDirty(false)
                }
                break;
            case 'ticketPrice':
                if (!e.target.value || isNaN(e.target.value) || e.target.value <= 0) {
                    setTicketPriceDirty(true);
                    setTicketPriceError('Поле должно содержать положительное число');
                } else {
                    setTicketPriceDirty(false);
                    setTicketPriceError('');
                }
                break;
            case 'ticketCount':
                if (!e.target.value || isNaN(e.target.value) || e.target.value <= 0) {
                    setTicketCountDirty(true);
                    setTicketCountError('Поле должно содержать положительное число');
                } else {
                    setTicketCountDirty(false);
                    setTicketCountError('');
                }
                break;
            default:
                break;
        }
    };

    return (
        <form className='reg' onSubmit={submitHandler}>
            <h1>Создание мероприятия</h1>

            {(nameDirty && nameError) && <div style={{ color: 'red' }}>{nameError}</div>}
            <div className="input-box">
                <input
                    onChange={e => setName(e.target.value)}
                    value={name}
                    name='name'
                    type="text"
                    placeholder="Название мероприятия..."
                    onBlur={blurHandler}
                />
            </div>

            {(nameCatDirty && nameCatError) && <div style={{ color: 'red' }}>{nameCatError}</div>}
            <div className="input-box">
                <input
                    onChange={e => setNameCat(e.target.value)}
                    value={nameCat}
                    name='nameCat'
                    type="text"
                    placeholder="Категория мероприятия..."
                    onBlur={blurHandler}
                />
            </div>

            {(descriptionDirty && descriptionError) && <div style={{ color: 'red' }}>{descriptionError}</div>}
            <div className="input-box">
                <input
                    onChange={e => setDescription(e.target.value)}
                    value={description}
                    name='description'
                    type="text"
                    placeholder="Описание..."
                    onBlur={blurHandler}
                />
            </div>

            {(dateTimeDirty && dateTimeError) && <div style={{ color: 'red' }}>{dateTimeError}</div>}
            <div className="input-box">
                <input
                    onChange={e => setDateTime(e.target.value)}
                    value={dateTime}
                    name='dateTime'
                    placeholder="Год-Месяц-День"
                    onBlur={blurHandler}
                />
            </div>

            {(addressDirty && addressError) && <div style={{ color: 'red' }}>{addressError}</div>}
            <div className="input-box">
                <input
                    onChange={e => setAddress(e.target.value)}
                    value={address}
                    name='address'
                    type="text"
                    placeholder="Адрес..."
                    onBlur={blurHandler}
                />
            </div>

            {(ticketPriceDirty && ticketPriceError) && <div style={{ color: 'red' }}>{ticketPriceError}</div>}
            <div className="input-box">
                <input
                    onChange={e => setTicketPrice(e.target.value)}
                    value={ticketPrice}
                    name='ticketPrice'
                    type="text"
                    placeholder="Цена билета..."
                    onBlur={blurHandler}
                />
            </div>

            {(ticketCountDirty && ticketCountError) && <div style={{ color: 'red' }}>{ticketCountError}</div>}
            <div className="input-box">
                <input
                    onChange={e => setTicketCount(e.target.value)}
                    value={ticketCount}
                    name='ticketCount'
                    type="text"
                    placeholder="Количество билетов..."
                    onBlur={blurHandler}
                />
            </div>

            <div className="register-link">
                <p>Добавить изображение </p>
                <input type="file" ref={imageRef} />
            </div>

            <button className="register-link" type="submit">
                Создать
            </button>
        </form>
    );
};

export default CreateEvent;
