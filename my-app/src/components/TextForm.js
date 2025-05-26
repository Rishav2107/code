import React,{useState} from 'react'

//In React JSX, anything you put inside { ... } is treated as JavaScript.
export default function TextForm({heading}) {
    const handleOnChange = (event) => {
        //console.log("handle change");
        setText(event.target.value);
    }
    const handleUpClick = () => {
        //console.log("Uppercase was clicked");
        let newText = text.toUpperCase();
        setText(newText);
    }
    const handleLoClick = () => {
        //console.log("Lowercase was clicked");
        let newText = text.toLocaleLowerCase();
        setText(newText);
    }
  const[text,setText] = useState('Enter text here');
  return (
    <>
    <div className="container">
        <h1>{heading}</h1>
        <div className="mb-3">
        <textarea className="form-control" value={text} onChange={handleOnChange} id="myBox" rows="8"></textarea>
        </div>
        <button className="btn btn-primary mx-2" onClick={handleUpClick}>Convert to UpperCase</button>
        <button className="btn btn-primary mx-2" onClick={handleLoClick}>Convert to LowerCase</button>
    </div>
    <div className="container my-3">
        <h1>Your text summary:</h1>
            <p>
                {text.split(" ").length} words , {text.length} characters
            </p>
    </div>
    </>
  )
}
