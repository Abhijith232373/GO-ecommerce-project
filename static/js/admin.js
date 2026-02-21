function openModal(){
    document.getElementById("userModal").style.display = "flex";
}

function closeModal(){
    document.getElementById("userModal").style.display = "none";
}

/* Close if click outside box */
window.onclick = function(event) {
    const modal = document.getElementById("userModal");
    if (event.target === modal) {
        modal.style.display = "none";
    }
}

document.getElementById("searchInput").addEventListener("keyup", function(){

let filter = this.value.toLowerCase()

let rows = document.querySelectorAll("#usersTable tbody tr")

rows.forEach(row=>{
let text = row.innerText.toLowerCase()
row.style.display = text.includes(filter) ? "" : "none"
})

})

function showAddUser() {
    document.getElementById("addUserForm").style.display = "block";
}

function hideAddUser() {
    document.getElementById("addUserForm").style.display = "none";
}

// ////////////Product ////////////
function openProductModal(){

document.getElementById("productModal").style.display="flex"

}

function closeProductModal(){

document.getElementById("productModal").style.display="none"

}