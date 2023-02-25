import React from "react"
import ChatInput from "../../components/atoms/ChatInput"
import ChatHeader from "../../components/organisms/chat/ChatHeader"
import ChatSection from "../../components/organisms/chat/ChatSection"
import { MessageContextProvider } from "../../contexts/MessageContext"

export default function Chat() {
  return (
    <MessageContextProvider>
      <ChatHeader />
      <ChatSection />
      <ChatInput />
    </MessageContextProvider>
  )
}
