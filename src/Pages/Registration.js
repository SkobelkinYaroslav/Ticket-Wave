import { useEffect, useState } from "react";
import '../App.css';
import axios from "axios";
const Registration=()=>{
    const[firstName, setfirstName] =useState('')
    const[lastName, setlastName] =useState('')
    const[email, setEmail] =useState('')
    const[password, setPassword] =useState('')

    const[firstNameDirty, setfirstNameDirty] =useState(false)
    const[lastNameDirty, setlastNameDirty] =useState(false)
    const[emailDirty, setEmailDirty] =useState(false)
    const[passwordDirty, setPasswordDirty] =useState(false)
    const[formValid, setFormValid] = useState(false)
    const[chooseRole, setchooseRole ]=useState(false)

    const[firstNameError, setfirstNameError] = useState('Поле не может быть пустым')
    const[lastNameError, setlastNameError] = useState('Поле не может быть пустым')
    const[emailError, setEmailError] = useState('Поле не может быть пустым')
    const[passwordError, setPasswordError] = useState('Поле не может быть пустым')
    const [selected, setSelected] = useState("Выберите роль")
    const optionsrole =['Пользователь', 'Организатор']

   
    const handleRoleSelection = (option) => {
        setSelected(option);
        setchooseRole(false);
    }

    const handleRegistration = () => {
        if (selected !== "Выберите роль"){
            const userData = {
                firstName: firstName,
                lastName: lastName,
                email: email,
                password: password,
                role: ''
            };

            if (selected === 'Пользователь') {
                userData.role = 'user';
            } else if (selected === 'Организатор') {
                userData.role = 'organizer';
            }

            axios.post('http://localhost:8080/register', userData)
                .then(function (response) {
                    window.location.href = '/log';
                })
                .catch(function (error) {
                    console.log('Error:', error);
                });
        }
    }



    useEffect(()=>{
        if(emailError || passwordError || firstNameError || lastNameError)
        {
            setFormValid(false)

        } else{
            setFormValid(true)
        }
    }, [emailError, passwordError, firstNameError, lastNameError])

    const firstNameHandler =(e) =>{
        setfirstName(e.target.value)
        
            if (!e.target.value)
            {
                setfirstNameError('Поле не может быть пустым')
            }
        else{
            setfirstNameError("")
        }
    }
    const lastNameHandler =(e) =>{
        setlastName(e.target.value)
        
            if (!e.target.value)
            {
                setlastNameError('Поле не может быть пустым')
            }
        else{
            setlastNameError("")
        }
    }
    const emailHandler=(e) =>{
        setEmail(e.target.value)
        const re =/^(([^<>()[\].,;:\s@"]+(\.[^<>()[\].,;:\s@"]+)*)|(".+"))@(([^<>()[\].,;:\s@"]+\.)+[^<>()[\].,;:\s@"]{2,})$/iu;
         
        if(!re.test(String(e.target.value).toLowerCase()))
        {
            setEmailError('Введен некорректный email')
        }
        else{
            setEmailError("")
        }
    }
    const passwordHandler =(e) =>{
        setPassword(e.target.value)
        if(e.target.value.length<3 || e.target.value.length>8 )
        {
            setPasswordError("Пароль должен быть не короче 3 и не длиннее 8 символов")
            if (!e.target.value)
            {
                setPasswordError('Поле не может быть пустым')
            }
        }
        else{
            setPasswordError("")
        }
    }
    const blurHandler = (e )=>{
       
        switch(e.target.name)
        {
            case 'firstName':
                setfirstNameDirty(true)
                break
            case 'lastName':
                    setlastNameDirty(true)
                    break
            case 'email':
                setEmailDirty(true)
                break
            case 'password':
                setPasswordDirty(true)
                break
            default:
                break
                
        }
    }
return (
    <div className='reg'>
       
            <h1>Регистрация</h1>
            <div className="chooserole">
                <div className="chooserole-btn" onClick={() => setchooseRole(!chooseRole)}>{selected}</div>
                {chooseRole && (
                    <div className="chooserole-content">
                        {optionsrole.map(option => (
                            <div key={option} onClick={() => handleRoleSelection(option)} className="chooserole-item">{option}</div>
                        ))}
                    </div>
                )}
            </div>
            {/* <div className="chooserole">
            <div className="chooserole-btn" onClick={e =>setchooseRole(!chooseRole)}>{selected}</div>
            {chooseRole &&( <div className="chooserole-content">{optionsrole.map(option => (<div onClick={(e) =>{setSelected(option)
            setchooseRole(false)}} className="chooserole-item">{option}</div>))}</div>)}</div> */}

            {(firstNameDirty && firstNameError ) && <div style={{color:'red'}}>{firstNameError}</div>}
            <div className="input-box"><input onChange={e => firstNameHandler(e)} value={firstName} onBlur={e=>blurHandler(e) } name='firstName' type="text" placeholder="Введите ваше имя..."/></div>
            {(lastNameDirty && lastNameError ) && <div style={{color:'red'}}>{lastNameError}</div>}
            <div className="input-box"><input onChange={e => lastNameHandler(e)} value={lastName} onBlur={e=>blurHandler(e) } name='lastName' type="text" placeholder="Введите вашу фамилию..."/></div>
            {(emailDirty && emailError ) && <div style={{color:'red'}}>{emailError}</div>}
            <div className="input-box"><input onChange={e => emailHandler(e)} value={email} onBlur={e=>blurHandler(e) } name='email' type="text" placeholder="Введите ваш email..."/></div>
            {(passwordDirty && passwordError) && <div style={{color:'red'}}>{passwordError}</div>}
            <div className="input-box"><input onChange={e => passwordHandler(e)} value={password} onBlur={e=>blurHandler(e) } name='password' type="password" placeholder="Введите ваш пароль..."/></div>

            {/* <div ><button type="submit" disabled={!formValid} ><p><a href="/mainpage">Зарегистрироваться</a></p></button> </div> */}
             {/* <div ><button className="register-link" disabled={!formValid} ><p><a href="/mainpage">Зарегистрироваться</a></p></button> </div> */}

           <div> <button className="register-link"  onClick={handleRegistration}>
                <p><a>Зарегистрироваться</a></p>
            </button></div> 
             
            {/* <div className="register-link"><p>Уже есть аккаунт? <a href="#">Войти</a></p></div> */}
            <div className="register-link"><p>Уже есть аккаунт? <a href="/log">Войти</a></p></div>

    </div>
);
}

export default Registration;




   



    

   

 