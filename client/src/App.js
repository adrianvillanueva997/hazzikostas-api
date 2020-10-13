import React, {Component} from 'react';
import './App.css';

class App extends Component {

    constructor() {
        super(undefined);
        this.state = {characters: []}
    }

    componentDidMount() {
        fetch("/api/v1/characters")
            .then((response) => response.json())
            .then((data) => {
                this.setState({characters: data})
            })
    }

    render() {
        return (
            <div>
                <h1 id='title'> Hazzikostas tracked toons</h1>
                <table id='characters' class="table table-striped">
                    <th>Toon</th>
                    <th>Region</th>
                    <th>Realm</th>
                    <th>All</th>
                    <th>DPS</th>
                    <th>Healer</th>
                    <th>Tank</th>
                    <th>Spec 0</th>
                    <th>Spec 1</th>
                    <th>Spec 2</th>
                    <th>Spec 3</th>
                    <th>Rank Overall</th>
                    <th>Rank Class</th>
                    <th>Rank Faction</th>
                    <tbody>
                    {
                        this.state.characters.map((character, index) => {
                            const {
                                toon_name, region, realm, all, dps,
                                healer, tank, spec_0, spec_1, spec_2, spec_3,
                                rank_overall, rank_class, rank_faction
                            } = character
                            return <tr>
                                <td>{toon_name}</td>
                                <td>{region}</td>
                                <td>{realm}</td>
                                <td>{all}</td>
                                <td>{dps}</td>
                                <td>{healer}</td>
                                <td>{tank}</td>
                                <td>{spec_0}</td>
                                <td>{spec_1}</td>
                                <td>{spec_2}</td>
                                <td>{spec_3}</td>
                                <td>{rank_overall}</td>
                                <td>{rank_class}</td>
                                <td>{rank_faction}</td>
                            </tr>
                        })}
                    </tbody>
                </table>
            </div>
        )
    }
}


export default App