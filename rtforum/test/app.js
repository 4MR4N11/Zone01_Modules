async function test() {
  const res = await fetch("http://localhost:8000/api/login", {
    method: "GET",
  });
  const data = await res.text();
  console.log(data);
}

test();
