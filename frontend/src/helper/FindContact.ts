type Contact = {
    ID: number,
    Firstname: string,
    Lastname: string
}

export const FindContact = (id: number, contacts: Contact[]) => {
    let contactFound: Contact = {ID: 0, Firstname: "", Lastname:""}
    contacts.forEach((contact: Contact) => {
        if(contact.ID === id) {
            contactFound = contact
        } 
    })
    return contactFound
}