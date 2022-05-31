const url = "https://api.cloudinary.com/v1_1/eternalq/image/upload";
const input = document.getElementById("imgInput");
const output = document.getElementById("imgURL");
const voutput = document.getElementById("imgURLp");

input.addEventListener("change", (e) => {
  e.preventDefault();

  const files = document.querySelector("[type=file]").files;
  const formData = new FormData();

  let file = files[0];
  console.log(file);
  formData.append("file", file);
  formData.append("upload_preset", "vr4neowj");

  fetch(url, {
    method: "POST",
    body: formData,
  })
    .then((response) => {
      console.log(response);
      return response.text();
    })
    .then((data) => {
      let obj = JSON.parse(data);
      output.value = obj.url;
      voutput.value = obj.url;
    });
});
