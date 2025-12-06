import {useEffect, useState} from 'react'

export default function Home(){
  const [health, setHealth] = useState(null)
  const [items, setItems] = useState([])

  useEffect(()=>{
    fetch(process.env.NEXT_PUBLIC_API_BASE_URL ? `${process.env.NEXT_PUBLIC_API_BASE_URL}/health` : 'http://localhost:8000/health')
      .then(r=>r.json())
      .then(setHealth)
      .catch(e=>setHealth({error: String(e)}))

    fetch(process.env.NEXT_PUBLIC_API_BASE_URL ? `${process.env.NEXT_PUBLIC_API_BASE_URL}/items` : 'http://localhost:8000/items')
      .then(r=>r.json())
      .then(setItems)
      .catch(e=>console.error(e))
  },[])

  return (
    <main style={{padding:24,fontFamily:'system-ui,Segoe UI,Roboto'}}>
      <h1>3-Tier App (Next.js)</h1>
      <section>
        <h2>API Health</h2>
        <pre>{JSON.stringify(health,null,2)}</pre>
      </section>
      <section>
        <h2>Items</h2>
        <ul>
          {items.map(it=> <li key={it.id}>{it.name} (#{it.id})</li>)}
        </ul>
      </section>
    </main>
  )
}
