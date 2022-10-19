import { SetStateAction, useEffect, useState } from "react";



function Apppppp() {
      const [arrayA, setA] = useState([1, 2]);
      const [objB, setB] =useState(new Map([]));
      useEffect(() => {
        arrayA.forEach((num) => {
          setB((oldValue: { set: (arg0: number, arg1: number) => any; }) => oldValue.set(num, num * num));
        });
      }, []);
    
      return <div>null</div>;
}
export default Apppppp;