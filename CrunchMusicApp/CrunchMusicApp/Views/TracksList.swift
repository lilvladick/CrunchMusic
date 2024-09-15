import SwiftUI

struct TracksList: View {
    @State private var navigationPath: [String] = []

    var body: some View {
        NavigationStack(path: $navigationPath) {
            VStack {
                // JSON PARSE TRACKS TO LIST
            }
            .toolbar {
                ToolbarItem(placement: .topBarLeading, content: {
                    Button(action: {
                        navigationPath.append("menu")
                    }, label: {
                        Image(systemName: "list.dash")
                    })
                })
                ToolbarItem(placement: .topBarTrailing, content: {
                    Button(action: {
                        // Action for the magnifying glass button
                    }, label: {
                        Image(systemName: "magnifyingglass")
                    })
                })
            }
            .bold()
            .navigationDestination(for: String.self) { destination in
                if destination == "menu" {
                    SideMenuView()
                }
            }
        }
    }
}

#Preview {
    TracksList()
}
