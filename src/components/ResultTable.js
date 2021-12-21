import './App.css';
import React from 'react';

class ResultTable extends React.Component {
    render() {
        if (this.props.items) {
            var items = this.props.items;
        }
        else {
            var items = [];
        }
        return (<table>
            <tr className="table_id">
                <th>Id</th>
                <th>Book Title</th>
                <th>Author</th>
                <th>ISBN</th>
                <th>Owner</th>
            </tr>
            {items.map(item => (
                <tr>
                    <td className="table_id">{item.id}</td>
                    <td>{item.title}</td>
                    <td>{item.author}</td>
                    <td>{item.isbn}</td>
                    <td>{item.owner}</td>
                </tr>
            ))}
        </table>);
    }
}

export default ResultTable;