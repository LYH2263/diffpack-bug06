async function load() {
  const h = await fetch("/api/health").then(r => r.json());
  document.getElementById("health").textContent = JSON.stringify(h);
  const jobs = await fetch("/api/jobs").then(r => r.json());
  const ul = document.getElementById("jobs");
  ul.innerHTML = "";
  jobs.forEach(j => {
    const li = document.createElement("li");
    li.textContent = j.ID + " " + j.Status;
    li.onclick = async () => {
      const hunks = await fetch("/api/jobs/" + j.ID + "?action=hunks").then(r => r.json());
      document.getElementById("detail").textContent = JSON.stringify(hunks, null, 2);
    };
    ul.appendChild(li);
  });
}
document.getElementById("refresh").onclick = load;
load();
