import axiosInstance from "@/services/axios";

// Invia un messaggio in una chat
export async function sendMessage(chatId, content) {
  try {
    console.log("Funzione SendMessage chiamata")
    await axiosInstance.post(`/chats/${chatId}/messages`, { content });
    console.log("Messaggio inviato con successo");
  } catch (error) {
    console.error("Errore durante l'invio del messaggio:", error);
    throw error;
  }
}

// Cancella una chat
export async function deleteChat(chatId) {
  try {
    await axiosInstance.delete(`/chats/${chatId}`);
    console.log("Chat eliminata con successo");
  } catch (error) {
    console.error("Errore durante l'eliminazione della chat:", error);
    throw error;
  }
}

// Crea una nuova chat privata
export async function createPrivateChat(user) {
  try {
    console.log("User della funzione createPrivateChat: " + user)
    const res = await axiosInstance.post(`/chats`, {
      name: user.name,
      users: [user.userId],
      isGroup: false
    });
    console.log("Chat privata creata con successo");
    return res.data.chatId; // ritorniamo l'ID della chat creata
  } catch (error) {
    console.error("Errore durante la creazione della chat:", error);
    throw error;
  }
}
