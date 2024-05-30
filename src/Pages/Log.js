import { useEffect, useState } from "react";
import '../App.css';
import axios from "axios";

const Log=()=> {
    
    const[email, setEmail] =useState('')
    const[password, setPassword] =useState('')
    
    const[emailDirty, setEmailDirty] =useState(false)
    const[passwordDirty, setPasswordDirty] =useState(false)
   
    const[emailError, setEmailError] =useState('Поле не может быть пустым')
    const[passwordError, setPasswordError] =useState('Поле не может быть пустым')
    // const[formValid, setFormValid] = useState(false)

    const[chooseRole, setchooseRole ]=useState(false)
    const[selected, setSelected] =useState("Выберите роль")
    const optionsrole =['Пользователь', 'Организатор']

    const handleRoleSelection = (option) => {
        setSelected(option);
        setchooseRole(false);
    }

    const handleRegistration = () => {
        const userData = {
            email: email,
            password: password,
        };


        axios.post('http://localhost:8080/login', userData,{
            withCredentials: true
        })
            .then(function (response) {
                console.log(response)

                if (response.data.message.role === 'user') {
                    window.location.href = '/mainpage';
                } else if (response.data.message.role === 'organizer'){
                    window.location.href = '/organizer-profile';
                }
            })
            .catch(function (error) {
                console.log('Error:', error);
            });
    }


    // useEffect(()=>{
    //     if(emailError || passwordError )
    //     {
    //         setFormValid(false)

    //     } else{
    //         setFormValid(true)
    //     }
    // }, [emailError, passwordError])

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
        
            <h1>Вход</h1>

            {/*<div className="chooserole">*/}
            {/*    <div className="chooserole-btn" onClick={() => setchooseRole(!chooseRole)}>{selected}</div>*/}
            {/*    {chooseRole && (*/}
            {/*        <div className="chooserole-content">*/}
            {/*            {optionsrole.map(option => (*/}
            {/*                <div key={option} onClick={() => handleRoleSelection(option)} className="chooserole-item">{option}</div>*/}
            {/*            ))}*/}
            {/*        </div>*/}
            {/*    )}*/}
            {/*</div>*/}
          
            {/* <div className="chooserole">
            <div className="chooserole-btn" onClick={e =>setchooseRole(!chooseRole)}>{selected}</div>
            {chooseRole &&( <div className="chooserole-content">{optionsrole.map(option => (<div onClick={(e) =>{setSelected(option)
            setchooseRole(false)}} className="chooserole-item">{option}</div>))}</div>)}</div>
             */}
            {(emailDirty && emailError ) && <div style={{color:'red'}}>{emailError}</div>}
            <div className="input-box"><input onChange={e => emailHandler(e)} value={email} onBlur={e=>blurHandler(e) } name='email' type="text" placeholder="Введите ваш email..."/></div>
            {(passwordDirty && passwordError) && <div style={{color:'red'}}>{passwordError}</div>}
            <div className="input-box"><input onChange={e => passwordHandler(e)} value={password} onBlur={e=>blurHandler(e) } name='password' type="password" placeholder="Введите ваш пароль..."/></div>
            {/* <div ><button className="register-link" disabled={!formValid} ><p><a href="/mainpage">Войти</a></p></button> </div> */}
            <div> <button className="register-link" onClick={handleRegistration}>
                <p><a>Войти</a></p>
            </button></div> 
            <div className="register-link"><p>Еще нет аккаунта? <a href="/">Регистрация</a></p></div>
    </div>
);
};

export default Log;


