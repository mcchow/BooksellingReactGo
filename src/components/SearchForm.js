import './App.css';
import React from 'react';
import ResultTable from './ResultTable';

class SearchForm extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            error: null,
            isLoaded: false,
            items: [],
            value: "",
            author: "",
            isbn: "",
            owner: ""
        };
        this.handleChange = this.handleChange.bind(this);
        this.handleButtonClicked = this.handleButtonClicked.bind(this);
        this.handleAddBookClicked = this.handleAddBookClicked.bind(this);
        this.handleOwnerBookSearchClicked = this.handleOwnerBookSearchClicked.bind(this);
        this.handleAuthorChange = this.handleAuthorChange.bind(this);
        this.handleISBNChange = this.handleISBNChange.bind(this);
        this.handleOwnerChange = this.handleOwnerChange.bind(this);
    }

    handleButtonClicked(event) {
        if (this.state.value != "") {
            fetch("/posts/" + this.state.value)
                .then(res => res.json())
                .then(
                    (result) => {
                        this.setState({
                            isLoaded: true,
                            items: result
                        });
                    },
                    // Note: it's important to handle errors here
                    // instead of a catch() block so that we don't swallow
                    // exceptions from actual bugs in components.
                    (error) => {
                        this.setState({
                            isLoaded: true,
                            error
                        });
                    }
                )
            alert('A name was submitted: ' + this.state.value);
        }
        else {
            alert('please input a name');
        }
        event.preventDefault();
    }

    handleOwnerBookSearchClicked(event) {
        if (this.state.value != "" && this.state.owner != "") {
            fetch("/getbook/" + this.state.value + "/" + this.state.owner)
                .then(res => res.json())
                .then(
                    (result) => {
                        this.setState({
                            isLoaded: true,
                            items: result
                        });
                    },
                    // Note: it's important to handle errors here
                    // instead of a catch() block so that we don't swallow
                    // exceptions from actual bugs in components.
                    (error) => {
                        this.setState({
                            isLoaded: true,
                            error
                        });
                    }
                )
            alert('A name was submitted: ' + this.state.value);
        }
        else {
            alert('please input a name');
        }
        event.preventDefault();
    }

    handleAddBookClicked(event) {
        if (this.state.value != "" && this.state.author != "" && this.state.isbn != "" && this.state.owner != "") {
            const requestOptions = {
                method: 'POST',
            };
            fetch("/add/" + this.state.value + "/" + this.state.author + "/" + this.state.isbn + "/" + this.state.owner, requestOptions)
                .then(res => res.json())
                .then(message => alert(message))
        }
        else {
            alert('please input all value');
        }
        event.preventDefault();
    }


    handleChange(event) {
        this.setState({ value: event.target.value });
    }

    handleAuthorChange(event) {
        this.setState({ author: event.target.value });
    }

    handleISBNChange(event) {
        this.setState({ isbn: event.target.value });
    }

    handleOwnerChange(event) {
        this.setState({ owner: event.target.value });
    }

    render() {
        return (<form onSubmit={this.handleSubmit}>
            <label>
                BookName:
                <input type="text" value={this.state.value} onChange={this.handleChange} maxLength="100" />
            </label>
            <label>
                Author:
                <input type="text" value={this.state.author} onChange={this.handleAuthorChange} maxLength="100" />
            </label>
            <label>
                ISBN:
                <input type="text" value={this.state.isbn} onChange={this.handleISBNChange} maxLength="13" />
            </label>
            <label>
                Owner:
                <input type="text" value={this.state.owner} onChange={this.handleOwnerChange} maxLength="100" />
            </label>
            <br />
            <input type="submit" value="Search By Bookname" onClick={this.handleButtonClicked.bind(this)} />
            <input type="submit" value="Search By Bookname and Owner" onClick={this.handleOwnerBookSearchClicked.bind(this)} />
            <input type="submit" value="Add book" onClick={this.handleAddBookClicked.bind(this)} />
            <ResultTable items={this.state.items} />
        </form>);
    }
}

export default SearchForm;