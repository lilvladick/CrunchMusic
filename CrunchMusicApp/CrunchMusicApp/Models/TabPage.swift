enum TabPage: String, CaseIterable {
    case music = "Music"
    case playlist = "Playlist"
    case settings = "Settings"
    
    var icon: String {
        switch self {
        case .music: return "music.note"
        case .playlist: return "music.note.list"
        case .settings: return "gearshape.fill"
        }
    }
    
    var label: String {
        self.rawValue
    }
}
