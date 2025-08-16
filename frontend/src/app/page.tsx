import createApolloClient from "@/apollo-client";
import { gql } from "@apollo/client";

export default async function Home() {
  const client = createApolloClient();
  const { data } = await client.query({
    query: gql`
      query TodoQuery($todoId: ID!) {
        todo(id: $todoId) {
          id
          text
          done
          user {
            id
          }
          createdAt
          updatedAt
        }
      }
    `,
    variables: {
      todoId: "1",
    },
  });
  return (
    <div className="font-sans grid grid-rows-[20px_1fr_20px] items-center justify-items-center min-h-screen p-8 pb-20 gap-16 sm:p-20">
      Test
      <pre>{JSON.stringify(data.todo, null, 2)}</pre>
    </div>
  );
}
