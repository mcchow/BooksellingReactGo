
import './App.css';
import React from 'react';
import SearchForm from './components/SearchForm';

class App extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      error: null,
      isLoaded: false,
      items: []
    };
  }

  componentDidMount() {
    /*
    fetch("/posts")
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
      )*/
  }

  render() {
    const { error, isLoaded, items } = this.state;
    return (
      <div>
        <SearchForm items={this.state.items} />
      </div>)
  }
}

export default App;
