import { useState } from 'react'
import './App.css'
import axios from 'axios'


axios.get('https://randomuser.me/api/?results=125')
  .then(function (res) {
    // handle success
    console.log('response:',res.data.results[0].name.first);
    console.log('response:',res.data.results[0].picture.large);
  })
  .catch(function (error) {
    // handle error
    console.log(error);
  })
  .finally(function () {
    // always executed
  });

function App() {
  const [zahl, setZahl] = useState(1)
  const [a, setA] = useState(27)

  return (
    <div className="AAA">
     
        <button onClick={()=> setZahl((zahl)=>zahl*23)}>
          count is {zahl}
        </button>
      <button onClick={()=>setA((a)=>a*27)}>
        zahl ist {a}
      </button>
      {a+ zahl}
    </div>
  )
}

export default App
