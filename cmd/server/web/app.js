document.getElementById('queryForm').addEventListener('submit', function (e) {
  e.preventDefault();

  const tenantId = document.getElementById('tenantId').value;
  const startDate = document.getElementById('startDate').value;
  const endDate = document.getElementById('endDate').value;

  let url = `http://localhost:8000/api/messages?start_date=${startDate}&end_date=${endDate}`;

  if (tenantId) {
    url += `&tenant_id=${tenantId}`;
  }

  fetch(url)
    .then(response => response.json())
    .then(data => {
      document.getElementById('result').innerText = JSON.stringify(data, null, 2);
    })
    .catch(error => {
      console.error('Error:', error);
    });
});
