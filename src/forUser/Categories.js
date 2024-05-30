import React from "react";
import { Component } from "react";

export class Categories extends Component{
    constructor(props){
        super(props)
        this.state={
            categories:[
                {
                    key: 'all',
                    nameCat: 'Все мероприятия'
                },
                {
                    key: 'concert',
                    nameCat: 'Концерт'
                },
                {
                    key: 'performance',
                    nameCat: 'Спектакль'
                },
                {
                    key: 'cinema',
                    nameCat: 'Кино'
                },
                {
                    key: 'standup',
                    nameCat: 'Стендап'
                }

            ]
        }
    }
    render(){
        return(
            <div className="categories">
                {this.state.categories.map(el=>(
                    <div key={el.key} onClick={()=>this.props.chooseCategory(el.key)}>{el.nameCat}</div>)
                )}
            </div>

        )
    }
}
export default Categories