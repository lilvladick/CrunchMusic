import SwiftUI

struct PlaylistsView: View {
    var body: some View {
        NavigationStack {
            VStack {
                //делаем запрос и по структуре создаем список плейлистов
                //с navigationLink и onDelete, которая будет посылать запрос на удаление
            }
            .navigationTitle("Your Playlists")
        }
    }
}

#Preview {
    PlaylistsView()
}
